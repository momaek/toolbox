package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryString(t *testing.T) {
	querys := NewQuerys()
	querys.Set("uid", 123)
	querys.Like("sign", "七牛%")
	querys.Between("created", "22", "33")
	querys.LessThan("count", 23)
	querys.GreaterThan("age", 11)
	querys.LessThanEqual("top", 99)
	querys.GreaterThanEqual("job", 123)
	querys.NotEqual("name", "bob")
	querys.Null("job")
	querys.NotNull("salary")

	assert.Equal(t, true, querys.CanSetToGormWhere())

	querys.OrderBy("uid", "desc")
	querys.Limit(10, 20)
	querys.In("id", []string{"1", "2", "3", "4", "5"})
	str := querys.String()

	assert.Equal(t, false, querys.CanSetToGormWhere())
	assert.Equal(t, "uid = 123 AND sign LIKE '七牛%' AND created BETWEEN '22' AND '33' AND count < 23 AND age > 11 AND top <= 99 AND job >= 123 AND name != 'bob' AND job IS NULL AND salary IS NOT NULL AND id IN ('1','2','3','4','5') ORDER BY uid desc LIMIT 10 OFFSET 20", str)

	q2 := NewQuerys()
	q2.Equal("uid", "222")
	q2.Limit(1, 0)
	assert.Equal(t, "uid = '222' LIMIT 1", q2.String())
}
