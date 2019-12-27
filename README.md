# findAge

HOW TO USE
reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input : ")
	text, _ := reader.ReadString('\n')
	inputs := strings.Split(text, ",")

	errs, result := findAge.Find(inputs)
	if errs != nil {
		for _, err := range errs {
			//print
			color.Red(err.Error())
		}
	} else {
		//print
		color.Green(result)
	}
