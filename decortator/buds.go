package main

type typeC struct {
	samsungPhone SamsungP
}

func (c *typeC) getPrice() int {
	samsungPrice := c.samsungPhone.getPrice()
	return samsungPrice + 7
}
