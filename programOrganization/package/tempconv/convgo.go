package tempconv

func FToC(f Fashreheit) Celsius{
	return Celsius((f-32)*5/9)
}

func CToF(c Celsius) Fashreheit{
	return Fashreheit(c*9/5+32)
}