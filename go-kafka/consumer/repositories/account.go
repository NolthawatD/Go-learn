package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&BankAccount{})
	return accountRepository{db}
}

func (obj accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Save(bankAccount).Error
}

func (obj accountRepository) Delete(id string) error {
	return obj.db.Where("id=?", id).Delete(&BankAccount{}).Error
}

func (obj accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = obj.db.Find(&bankAccounts).Error
	return bankAccounts, err
}

func (obj accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	err = obj.db.Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}
