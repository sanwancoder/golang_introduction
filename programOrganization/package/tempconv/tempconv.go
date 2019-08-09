package tempconv

import "fmt"

type Celsius float64  //摄氏温度
type Fashreheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15  //绝对零度
	FreezingC Celsius = 0    //结冰点温度
	BoilingC Celsius = 100 		//沸水温度
)

func (c Celsius) String() string  {
	return fmt.Sprintf("%g °C",c)
}

func (f Fashreheit) String() string  {
	return fmt.Sprintf("%g °F",f)
}
