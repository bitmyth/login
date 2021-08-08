package orm

import "gorm.io/gorm"

type Gorm struct {
    DB *gorm.DB
}

func (g *Gorm) Where(query interface{}, args ...interface{}) (tx *Gorm) {
    return &Gorm{g.DB.Where(query)}
}

func (g *Gorm) Find(dest interface{}, conds ...interface{}) (tx *Gorm) {
    return &Gorm{g.DB.Find(dest)}
}

func (g *Gorm) First(dest interface{}, conds ...interface{}) (tx *Gorm) {
    return &Gorm{g.DB.First(dest)}
}

func (g *Gorm) Save(value interface{}) (tx *Gorm) {
    return &Gorm{g.DB.Save(value)}
}

func (g *Gorm) Error() error {
    return g.DB.Error
}
