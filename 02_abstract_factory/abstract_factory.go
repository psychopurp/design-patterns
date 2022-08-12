package abstractfactory

import "fmt"

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	Factory func(string) interface{}
)

func (mdb *mongoDB) GetData(query string) string {
	if val, ok := mdb.database[query]; ok {
		fmt.Println("MongoDB")
		return val
	} else {
		return ""
	}
}

func (mdb *mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql *sqlite) GetData(query string) string {
	if val, ok := sql.database[query]; ok {
		fmt.Println("sqlite")
		return val
	} else {
		return ""
	}
}

func (sql *sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (ntfs *ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext *ext4) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ext.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs *ntfs) FindFile(path string) file {
	if file, ok := ntfs.files[path]; ok {
		return file
	}
	return file{}
}

func (ext *ext4) FindFile(path string) file {
	if file, ok := ext.files[path]; ok {
		return file
	}
	return file{}
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return &mongoDB{database: make(map[string]string)}
	case "development":
		return &sqlite{database: make(map[string]string)}
	default:
		return nil
	}
}

func FileSystemFactory(env string) interface{} {
	switch env {
	case "production":
		return &ntfs{files: make(map[string]file)}
	case "development":
		return &ext4{files: make(map[string]file)}
	default:
		return nil
	}
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FileSystemFactory
	default:
		return nil
	}
}

func SetupConstrcutors(env string) (Database, FileSystem) {
	fs := AbstractFactory("filesystem")
	db := AbstractFactory("database")
	return db(env).(Database), fs(env).(FileSystem)
}
