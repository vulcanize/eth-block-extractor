package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"
)

type IpfsManagerMock struct {
	configEnsured bool
	ipfsPath      string
}

func (manager *IpfsManagerMock) EnsureConfig(ipfsPath string) error {
	manager.configEnsured = true
	manager.ipfsPath = ipfsPath
	return nil
}

type IpfsManagerWithErrorMock struct{}

func (manager *IpfsManagerWithErrorMock) EnsureConfig(ipfsPath string) error {
	return errors.New("An error configuring IPFS occurred")
}

type PostgresManagerMock struct {
	schemaEnsured  bool
	databaseConfig config.Database
}

func (manager *PostgresManagerMock) EnsureSchema(databaseConfig config.Database) error {
	manager.schemaEnsured = true
	manager.databaseConfig = databaseConfig
	return nil
}

type PostgresManagerWithErrorMock struct{}

func (manager *PostgresManagerWithErrorMock) EnsureSchema(databaseConfig config.Database) error {
	return errors.New("An error establishing Postgres schema occurred")
}

var _ = Describe("Init command successful operation", func() {
	var ipfsPath = "/an/example/ipfs/.path"

	It("Ensures presence of IPFS config", func() {
		var ipfsManagerMock = &IpfsManagerMock{configEnsured: false}
		var postgresManagerMock = &PostgresManagerMock{}
		var databaseConfig = config.Database{}

		initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(ipfsManagerMock.configEnsured).To(BeTrue())
	})

	It("Provides the IPFS path to the IPFS manager", func() {
		var ipfsManagerMock = &IpfsManagerMock{}
		var postgresManagerMock = &PostgresManagerMock{}
		var databaseConfig = config.Database{}

		initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(ipfsManagerMock.ipfsPath).To(Equal("/an/example/ipfs/.path"))
	})

	It("Ensures presence of Postgres schema for IPFS driver", func() {
		var ipfsManagerMock = &IpfsManagerMock{}
		var postgresManagerMock = &PostgresManagerMock{schemaEnsured: false}
		var databaseConfig = config.Database{}

		initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(postgresManagerMock.schemaEnsured).To(BeTrue())
	})

	It("Provides the database configuration to the Postgres manager", func() {
		var ipfsManagerMock = &IpfsManagerMock{}
		var postgresManagerMock = &PostgresManagerMock{schemaEnsured: false}
		var databaseConfig = config.Database{Hostname: "example.com"}

		initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(postgresManagerMock.databaseConfig).To(Equal(databaseConfig))
	})

	It("Returns an IPFS config error", func() {
		var ipfsManagerMock = &IpfsManagerWithErrorMock{}
		var postgresManagerMock = &PostgresManagerMock{schemaEnsured: false}
		var databaseConfig = config.Database{}

		var expectedErr = errors.New("An error configuring IPFS occurred")

		var initErr = initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(expectedErr.Error()).To(Equal(initErr.Error()))
	})

	It("Does not ensure schema when IPFS fails", func() {
		var ipfsManagerMock = &IpfsManagerWithErrorMock{}
		var postgresManagerMock = &PostgresManagerMock{schemaEnsured: false}
		var databaseConfig = config.Database{}

		initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(postgresManagerMock.schemaEnsured).To(BeFalse())
	})

	It("Returns a Postgres schema error", func() {
		var ipfsManagerMock = &IpfsManagerMock{}
		var postgresManagerMock = &PostgresManagerWithErrorMock{}
		var databaseConfig = config.Database{Hostname: "example.com"}

		var expectedErr = errors.New("An error establishing Postgres schema occurred")

		var initErr = initializeIpfsWithPostgres(ipfsManagerMock, ipfsPath, postgresManagerMock, databaseConfig)

		Expect(expectedErr.Error()).To(Equal(initErr.Error()))
	})
})
