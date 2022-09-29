package main

type buds struct {
	samsungPhone SamsungP
}

func (c *buds) getPrice() int {
	SamsungPrice := c.samsungPhone.getPrice()
	return SamsungPrice + 10
}
