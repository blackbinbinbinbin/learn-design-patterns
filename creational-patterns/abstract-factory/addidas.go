package main

// AddidasFactory 抽象方法
// 负责创建两个同系列（addidas）的产品，shoe 和 shirt
type AddidasFactory struct {
}

func (n AddidasFactory) makeShoe() IShoe {
	return &AddidasShoe{
		Shoe{
			logo: "Addidas",
			size: 9,
		},
	}
}

func (n AddidasFactory) makeShirt() IShirt {
	return &AddidasShirt{
		Shirt{
			logo: "Addidas",
			size: 19,
		},
	}
}
