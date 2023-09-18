package database


type Database struct{

	number int
	name []string

}


func (db *Database) construct(number int,name []string){
	db.name=name
	db.number=number

}

func (db Database) number() int{

	return db.number

}

func (db Database) number_of_names() int{

	return len(db.name)

}

func (db Database) get_name() []string{

	return db.name

}
