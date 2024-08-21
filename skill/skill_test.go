package skill

// func resetDB() {
// 	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
// 	defer db.Close()
// 	// h := Handler {Db: db}
// 	_, err := db.Exec("DROP TABLE if exists skill CASCADE;")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	_, err = db.Exec(`
// 		INSERT INTO skill (key, name, description, logo, tags)
// 		VALUES (
// 			'go',
// 			'Go',
// 			'Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.',
// 			'https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg',
// 			'{go, golang}'
// 		),
// 		(
// 			'nodejs',
// 			'Node.js',
// 			'Node.js is an open-source, cross-platform, JavaScript runtime environment that executes JavaScript code outside of a browser.',
// 			'https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg',
// 			'{runtime, javascript}')
// 			`)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
