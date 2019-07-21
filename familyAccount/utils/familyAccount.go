package utils
import (
	"fmt"
)

type FamilyAccount struct {
	key string
	loop bool
	balance float64
	money float64
	note string
	flag bool
	details string
}

//工厂模式构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		key:"",
		loop: true,
		balance:10000.0,
		money:0.0,
		note:"",
		flag:false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}


//展示明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录------------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支明细...")
	}
}

//将登记收入写成一个方法，和*FamilyAccount对象绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}

//将登记支出些写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次指出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n指出\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}

//将退出方法写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) exit() {
	fmt.Println("是否确定退出？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你输入有误")
	}
	if choice == "y"{
		this.loop = false
	}
}

//结构体绑定方法
func (this *FamilyAccount) MainMenu(){
	for{
		fmt.Println("------------------当其那收支明细记录----------")
		fmt.Println("				1 收支明细")
		fmt.Println("				2 登记收入")
		fmt.Println("				3 登记指出")
		fmt.Println("				4 推出软件")

		fmt.Print("请选择（1-4）")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()			
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}

	}
}