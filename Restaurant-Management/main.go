package main

func main(){
	port := os.Gentenv("PORT")
}

router := Gofr.new()
router.use(Gofr.Logger())
