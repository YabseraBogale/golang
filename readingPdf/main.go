package main

func main() {
<<<<<<< HEAD

=======
	file, err := os.ReadFile("simple.pdf")
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range file {
		fmt.Print(string(i))
	}
>>>>>>> 3d9e07e (ok)
}
