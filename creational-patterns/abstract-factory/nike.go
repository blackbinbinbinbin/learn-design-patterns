package main

// NikeFactory 抽象方法
// 负责创建两个同系列（nike）的产品，shoe 和 shirt
type NikeFactory struct {
}

func (n NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			logo: "Nike",
			size: 10,
		},
	}
}

func (n NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		Shirt{
			logo: "Nike",
			size: 20,
		},
	}
}
