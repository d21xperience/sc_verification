package config

import (
	"log"
	"sc-service/model"
	"time"

	"gorm.io/gorm"
)

func Seed() {
	networks := []model.BlockchainNetwork{
		{
			NetworkName:    "Ethereum Mainnet",
			BlockchainType: "Ethereum",
			Description:    "Main Ethereum network",
			Unit:           "ETH",
			RPCURL:         "https://mainnet.infura.io/v3/YOUR-PROJECT-ID",
			ChainID:        1,
		},
		{
			NetworkName:    "Quorum Testnet",
			BlockchainType: "Quorum",
			Description:    "Private Quorum network for testing",
			Unit:           "USD",
			RPCURL:         "http://localhost:8545",
			ChainID:        1337,
		},
	}

	for _, network := range networks {
		if err := DB.Create(&network).Error; err != nil {
			log.Fatalf("Failed to seed blockchain networks: %v", err)
		}
	}

	log.Println("Seeding completed!")
}

func SeedStudent() {
	// Students Data
	students := []model.Student{
		{
			StudentID:   1,
			Nama:        "John Doe",
			NIS:         "12345678",
			NISN:        "87654321",
			TptLahir:    time.Date(2005, 10, 12, 0, 0, 0, 0, time.UTC),
			TglLahir:    time.Date(2005, 10, 12, 0, 0, 0, 0, time.UTC),
			AsalSekolah: "SMA Negeri 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			StudentID:   2,
			Nama:        "Jane Smith",
			NIS:         "23456789",
			NISN:        "98765432",
			TptLahir:    time.Date(2006, 5, 21, 0, 0, 0, 0, time.UTC),
			TglLahir:    time.Date(2006, 5, 21, 0, 0, 0, 0, time.UTC),
			AsalSekolah: "SMA Negeri 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Certificates Data
	certificates := []model.Certificate{
		{
			CertificateID:   1,
			StudentID:       1,
			CertificateHash: "hash123...",
			IPFSCID:         "Qm12345...",
			IssueDate:       time.Now(),
			IssuerName:      "Ministry of Education",
			NilaiRata:       "85.5",
			BlockchainType:  "Ethereum",
			TransactionHash: "tx456...",
			ContractAddress: "0xcontract123...",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			CertificateID:   2,
			StudentID:       2,
			CertificateHash: "hash456...",
			IPFSCID:         "Qm67890...",
			IssueDate:       time.Now(),
			IssuerName:      "Ministry of Education",
			NilaiRata:       "90.0",
			BlockchainType:  "Hyperledger Fabric",
			TransactionHash: "tx789...",
			ContractAddress: "0xcontract456...",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	// Insert data into database
	err := DB.Transaction(func(tx *gorm.DB) error {
		// Seed Students
		if err := tx.Create(&students).Error; err != nil {
			return err
		}

		// Seed Certificates
		if err := tx.Create(&certificates).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	log.Println("Database seeded successfully.")
}

//func dbSeed() {
// 	// Dummy data (tambahkan jika database kosong)
// 	DB.Create(&model.BlockchainNetwork{
// 		NetworkName: "Ethereum-Local",
// 		Unit:        "ETH",
// 		Account: []model.Account{
// 			{Name: "Account 1", Address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 100, NetworkID: 1},
// 			{Name: "Account 2", Address: "9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 200, NetworkID: 1},
// 		},
// 	})

// 	DB.Create(&model.BlockchainNetwork{
// 		NetworkName: "Quorum",
// 		Unit:        "ETH",
// 		Account: []model.Account{
// 			{Name: "quo 1", Address: "quuro9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 500, NetworkID: 2},
// 			{Name: "quo 2", Address: "quuro2TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 120, NetworkID: 2},
// 		},
// 	})

// 	DB.Create(&model.BlockchainNetwork{
// 		NetworkName: "Hyperledger Fabric",
// 		Unit:        "No limit",
// 		Account: []model.Account{
// 			{Name: "HF-01", Address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 350, NetworkID: 3},
// 			{Name: "HF-02", Address: "10TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 130, NetworkID: 3},
// 		},
// 	})
// }
