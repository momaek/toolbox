package db

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type querier interface {
	String() string
	Index() int
}

type operator struct {
	Column string
	Value  interface{}
	Op     string
	idx    int
}

// String build operator string
func (e operator) String() string {
	ret := ""

	val := reflect.ValueOf(e.Value)
	if val.Kind() == reflect.String {
		ret = fmt.Sprintf("%s %s '%v'", e.Column, e.Op, e.Value)
	} else {
		ret = fmt.Sprint(e.Column, " ", e.Op, " ", e.Value)
	}

	return ret
}

// Index implement querier
func (e operator) Index() int {
	return e.idx
}

type between struct {
	Column string
	Start  interface{}
	End    interface{}
	idx    int
}

// String build between string
func (b between) String() string {
	ret := ""

	val := reflect.ValueOf(b.Start)
	if val.Kind() == reflect.String {
		ret = fmt.Sprintf("%s BETWEEN '%v' AND '%v'", b.Column, b.Start, b.End)
	} else {
		ret = fmt.Sprintf("%s BETWEEN %v AND %v", b.Column, b.Start, b.End)
	}

	return ret
}

// Index implement querier
func (b between) Index() int {
	return b.idx
}

type null struct {
	Column string
	IsNull bool
	idx    int
}

// String build null string
func (n null) String() string {
	if n.IsNull {
		return fmt.Sprintf("%s IS NULL", n.Column)
	}

	return fmt.Sprintf("%s IS NOT NULL", n.Column)
}

// Index implement querier
func (n null) Index() int {
	return n.idx
}

// Querys mysql query, donot support concurrency call
type Querys struct {
	queries map[int]querier
	count   int
	orderBy querier
	limit   querier
}

// NewQuerys new mysql querys
func NewQuerys() *Querys {
	return &Querys{
		queries: make(map[int]querier),
	}
}

// Set 设置 query 参数
func (q *Querys) Set(column string, value interface{}) {
	q.Equal(column, value)
}

func (q *Querys) setOperator(column, op string, value interface{}) {
	q.count++
	q.queries[q.count] = operator{
		Column: column,
		Value:  value,
		Op:     op,
		idx:    q.count,
	}
}

func (q Querys) whereBuf() *bytes.Buffer {
	buf := bytes.NewBuffer(nil)
	if len(q.queries) == 0 {
		return buf
	}

	queries := []querier{}
	for _, v := range q.queries {
		queries = append(queries, v)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i].Index() < queries[j].Index()
	})

	if len(queries) > 0 {
		buf.WriteString(queries[0].String())
		for _, v := range queries[1:] {
			buf.WriteString(" AND ")
			buf.WriteString(v.String())
		}
	}

	return buf
}

// CountString from mysql count, remove limit, order by
func (q Querys) CountString() string {
	return q.whereBuf().String()
}

// String 生成 mysql query where
func (q Querys) String() string {
	buf := q.whereBuf()
	if q.orderBy != nil {
		buf.WriteString(" ")
		buf.WriteString(q.orderBy.String())
	}

	if q.limit != nil {
		buf.WriteString(" ")
		buf.WriteString(q.limit.String())
	}

	return buf.String()
}

// Like  column like 'value'
func (q *Querys) Like(column, value string) *Querys {
	q.setOperator(column, "LIKE", value)
	return q
}

// LessThan column < value
func (q *Querys) LessThan(column string, value interface{}) *Querys {
	q.setOperator(column, "<", value)
	return q
}

// GreaterThan column > value
func (q *Querys) GreaterThan(column string, value interface{}) *Querys {
	q.setOperator(column, ">", value)
	return q
}

// LessThanEqual column <= value
func (q *Querys) LessThanEqual(column string, value interface{}) *Querys {
	q.setOperator(column, "<=", value)
	return q
}

// GreaterThanEqual column >= value
func (q *Querys) GreaterThanEqual(column string, value interface{}) *Querys {
	q.setOperator(column, ">=", value)
	return q
}

// NotEqual column != value
func (q *Querys) NotEqual(column string, value interface{}) *Querys {
	q.setOperator(column, "!=", value)
	return q
}

// Equal column = value
func (q *Querys) Equal(column string, value interface{}) *Querys {
	q.setOperator(column, "=", value)
	return q
}

// Between column Between start and end
func (q *Querys) Between(column string, start, end interface{}) *Querys {
	q.count++
	q.queries[q.count] = between{
		Column: column,
		Start:  start,
		End:    end,
		idx:    q.count,
	}

	return q
}

// Null column IS NULL
func (q *Querys) Null(column string) *Querys {
	q.count++
	q.queries[q.count] = null{column, true, q.count}
	return q
}

// NotNull column IS NOT NULL
func (q *Querys) NotNull(column string) *Querys {
	q.count++
	q.queries[q.count] = null{column, false, q.count}
	return q
}

// OrderBy ORDER BY column [opt]
// OrderBy 不支持直接放到 Gorm 的 Where 里面, 需要自己写 exec
func (q *Querys) OrderBy(column string, opt ...string) *Querys {
	o := orderBy{
		Column: column,
	}
	if len(opt) > 0 {
		o.Opt = opt[0]
	}

	q.orderBy = o
	return q
}

// Limit LIMIT [limit] OFFSET [offset]
// Limit 不支持直接放到 Gorm 的 Where 函数里面使用
func (q *Querys) Limit(limit, offset int) *Querys {
	q.limit = limitOffset{limit, offset}
	return q
}

// CanSetToGormWhere check if querys can used as Gorm Where
func (q *Querys) CanSetToGormWhere() bool {
	if q.orderBy != nil || q.limit != nil {
		return false
	}

	return true
}

// In  [column] IN (values...)
func (q *Querys) In(column string, values interface{}) *Querys {
	q.count++
	i := in{
		Column: column,
		idx:    q.count,
	}

	val := reflect.ValueOf(values)
	if val.Kind() == reflect.Slice {
		for ii := 0; ii < val.Len(); ii++ {
			i.Values = append(i.Values, val.Index(ii).Interface())
		}
	}

	q.queries[q.count] = i
	return q
}

type orderBy struct {
	Column string
	Opt    string // ASC, DESC
}

// String ORDER BY column [Opt]
func (o orderBy) String() string {
	s := fmt.Sprintf("ORDER BY %s", o.Column)
	if len(o.Opt) > 0 {
		s += fmt.Sprintf(" %s", o.Opt)
	}

	return s
}

// Index .
func (o orderBy) Index() int { return 0 }

type limitOffset struct {
	Limit  int
	Offset int
}

// Limit LIMIT [limit] OFFSET [offset]
func (l limitOffset) String() string {
	s := ""
	if l.Limit > 0 {
		s += fmt.Sprintf("LIMIT %d", l.Limit)
	}

	if l.Offset > 0 {
		s += fmt.Sprintf(" OFFSET %d", l.Offset)
	}

	return s
}

// Index ..
func (l limitOffset) Index() int { return 0 }

type in struct {
	Column string
	Values []interface{}
	idx    int
}

// String ..
func (i in) String() string {
	strArray := make([]string, len(i.Values))
	for k, v := range i.Values {

		val := reflect.ValueOf(v)
		if val.Kind() == reflect.String {
			strArray[k] = fmt.Sprintf("'%v'", v)
		} else {
			strArray[k] = fmt.Sprintf("%v", v)
		}
	}

	return fmt.Sprintf("%s IN (%s)", i.Column, strings.Join(strArray, ","))
}

// Index ..
func (i in) Index() int { return i.idx }
