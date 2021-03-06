// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: d
func (_m *DBClient) AddDevice(d models.Device) (models.Device, errors.EdgeX) {
	ret := _m.Called(d)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(models.Device) models.Device); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.Device) errors.EdgeX); ok {
		r1 = rf(d)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddDeviceProfile provides a mock function with given fields: e
func (_m *DBClient) AddDeviceProfile(e models.DeviceProfile) (models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) models.DeviceProfile); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.DeviceProfile) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddDeviceService provides a mock function with given fields: e
func (_m *DBClient) AddDeviceService(e models.DeviceService) (models.DeviceService, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.DeviceService
	if rf, ok := ret.Get(0).(func(models.DeviceService) models.DeviceService); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.DeviceService)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.DeviceService) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddProvisionWatcher provides a mock function with given fields: pw
func (_m *DBClient) AddProvisionWatcher(pw models.ProvisionWatcher) (models.ProvisionWatcher, errors.EdgeX) {
	ret := _m.Called(pw)

	var r0 models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) models.ProvisionWatcher); ok {
		r0 = rf(pw)
	} else {
		r0 = ret.Get(0).(models.ProvisionWatcher)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.ProvisionWatcher) errors.EdgeX); ok {
		r1 = rf(pw)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDeviceProfiles provides a mock function with given fields: offset, limit, labels
func (_m *DBClient) AllDeviceProfiles(offset int, limit int, labels []string) ([]models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(offset, limit, labels)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(int, int, []string) []models.DeviceProfile); ok {
		r0 = rf(offset, limit, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDeviceServices provides a mock function with given fields: offset, limit, labels
func (_m *DBClient) AllDeviceServices(offset int, limit int, labels []string) ([]models.DeviceService, errors.EdgeX) {
	ret := _m.Called(offset, limit, labels)

	var r0 []models.DeviceService
	if rf, ok := ret.Get(0).(func(int, int, []string) []models.DeviceService); ok {
		r0 = rf(offset, limit, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceService)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDevices provides a mock function with given fields: offset, limit, labels
func (_m *DBClient) AllDevices(offset int, limit int, labels []string) ([]models.Device, errors.EdgeX) {
	ret := _m.Called(offset, limit, labels)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(int, int, []string) []models.Device); ok {
		r0 = rf(offset, limit, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllProvisionWatchers provides a mock function with given fields: offset, limit, labels
func (_m *DBClient) AllProvisionWatchers(offset int, limit int, labels []string) ([]models.ProvisionWatcher, errors.EdgeX) {
	ret := _m.Called(offset, limit, labels)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(int, int, []string) []models.ProvisionWatcher); ok {
		r0 = rf(offset, limit, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteDeviceById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteDeviceByName provides a mock function with given fields: name
func (_m *DBClient) DeleteDeviceByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteDeviceProfileById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceProfileById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteDeviceProfileByName provides a mock function with given fields: name
func (_m *DBClient) DeleteDeviceProfileByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteDeviceServiceById provides a mock function with given fields: id
func (_m *DBClient) DeleteDeviceServiceById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteDeviceServiceByName provides a mock function with given fields: name
func (_m *DBClient) DeleteDeviceServiceByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeviceById provides a mock function with given fields: id
func (_m *DBClient) DeviceById(id string) (models.Device, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceByName provides a mock function with given fields: name
func (_m *DBClient) DeviceByName(name string) (models.Device, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceIdExists provides a mock function with given fields: id
func (_m *DBClient) DeviceIdExists(id string) (bool, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceNameExists provides a mock function with given fields: id
func (_m *DBClient) DeviceNameExists(id string) (bool, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfileByName provides a mock function with given fields: name
func (_m *DBClient) DeviceProfileByName(name string) (models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfileNameExists provides a mock function with given fields: name
func (_m *DBClient) DeviceProfileNameExists(name string) (bool, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByManufacturer provides a mock function with given fields: offset, limit, manufacturer
func (_m *DBClient) DeviceProfilesByManufacturer(offset int, limit int, manufacturer string) ([]models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(offset, limit, manufacturer)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(int, int, string) []models.DeviceProfile); ok {
		r0 = rf(offset, limit, manufacturer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, manufacturer)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByManufacturerAndModel provides a mock function with given fields: offset, limit, manufacturer, model
func (_m *DBClient) DeviceProfilesByManufacturerAndModel(offset int, limit int, manufacturer string, model string) ([]models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(offset, limit, manufacturer, model)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(int, int, string, string) []models.DeviceProfile); ok {
		r0 = rf(offset, limit, manufacturer, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, manufacturer, model)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByModel provides a mock function with given fields: offset, limit, model
func (_m *DBClient) DeviceProfilesByModel(offset int, limit int, model string) ([]models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(offset, limit, model)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(int, int, string) []models.DeviceProfile); ok {
		r0 = rf(offset, limit, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, model)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceServiceById provides a mock function with given fields: id
func (_m *DBClient) DeviceServiceById(id string) (models.DeviceService, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.DeviceService
	if rf, ok := ret.Get(0).(func(string) models.DeviceService); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.DeviceService)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceServiceByName provides a mock function with given fields: name
func (_m *DBClient) DeviceServiceByName(name string) (models.DeviceService, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.DeviceService
	if rf, ok := ret.Get(0).(func(string) models.DeviceService); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.DeviceService)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceServiceNameExists provides a mock function with given fields: name
func (_m *DBClient) DeviceServiceNameExists(name string) (bool, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DevicesByProfileName provides a mock function with given fields: offset, limit, profileName
func (_m *DBClient) DevicesByProfileName(offset int, limit int, profileName string) ([]models.Device, errors.EdgeX) {
	ret := _m.Called(offset, limit, profileName)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Device); ok {
		r0 = rf(offset, limit, profileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, profileName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DevicesByServiceName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) DevicesByServiceName(offset int, limit int, name string) ([]models.Device, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Device); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatcherByName provides a mock function with given fields: name
func (_m *DBClient) ProvisionWatcherByName(name string) (models.ProvisionWatcher, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(string) models.ProvisionWatcher); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.ProvisionWatcher)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatchersByProfileName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) ProvisionWatchersByProfileName(offset int, limit int, name string) ([]models.ProvisionWatcher, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(int, int, string) []models.ProvisionWatcher); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatchersByServiceName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) ProvisionWatchersByServiceName(offset int, limit int, name string) ([]models.ProvisionWatcher, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func(int, int, string) []models.ProvisionWatcher); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateDeviceProfile provides a mock function with given fields: e
func (_m *DBClient) UpdateDeviceProfile(e models.DeviceProfile) errors.EdgeX {
	ret := _m.Called(e)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) errors.EdgeX); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}
