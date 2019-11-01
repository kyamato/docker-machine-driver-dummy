package dummy

import (
	"errors"
	"fmt"
	"net"
	"os"
	// "path"
	"strconv"
	"time"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/mcnutils"
	"github.com/docker/machine/libmachine/state"
)

type Driver struct {
	*drivers.BaseDriver
	EnginePort int
	SSHKey     string
}

const (
	driverName = "dummy"
	defaultTimeout = 15 * time.Second
)

// GetCreateFlags registers the flags this driver adds to
// "docker hosts create"
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
//		mcnflag.IntFlag{
//			Name:   "dummy-engine-port",
//			Usage:  "Docker engine port",
//			Value:  engine.DefaultPort,
//			EnvVar: "GENERIC_ENGINE_PORT",
//		},
		mcnflag.StringFlag{
			Name:   "dummy-ip-address",
			Usage:  "IP Address of machine",
			EnvVar: "GENERIC_IP_ADDRESS",
		},
//		mcnflag.StringFlag{
//			Name:   "dummy-ssh-user",
//			Usage:  "SSH user",
//			Value:  drivers.DefaultSSHUser,
//			EnvVar: "GENERIC_SSH_USER",
//		},
//		mcnflag.StringFlag{
//			Name:   "dummy-ssh-key",
//			Usage:  "SSH private key path (if not provided, default SSH key will be used)",
//			Value:  "",
//			EnvVar: "GENERIC_SSH_KEY",
//		},
//		mcnflag.IntFlag{
//			Name:   "dummy-ssh-port",
//			Usage:  "SSH port",
//			Value:  drivers.DefaultSSHPort,
//			EnvVar: "GENERIC_SSH_PORT",
//		},
	}
}

// NewDriver creates and returns a new instance of the driver
func NewDriver(hostName, storePath string) drivers.Driver {
	//log.Info("func newdriver")
	return &Driver{
		EnginePort: engine.DefaultPort,
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	log.Info("func DriverName()")
	return driverName
}

func (d *Driver) GetSSHHostname() (string, error) {
	log.Info("func GetSSHHostname()")
	return d.GetIP()
}

func (d *Driver) GetSSHUsername() string {
	log.Info("func GetSSHUsername()")
	return d.SSHUser
}

func (d *Driver) GetSSHKeyPath() string {
	log.Info("func GetSSHKeyPass()")
	return d.SSHKeyPath
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
//	d.EnginePort = flags.Int("dummy-engine-port")
//	d.IPAddress = flags.String("dummy-ip-address")
//	d.SSHUser = flags.String("dummy-ssh-user")
//	d.SSHKey = flags.String("dummy-ssh-key")
//	d.SSHPort = flags.Int("dummy-ssh-port")

	d.EnginePort = 2376
	d.IPAddress = "127.0.0.1"
	d.SSHUser = "root"
	d.SSHKey = ""
	d.SSHPort = 22

	log.Info("func SetConfigFromFlags()")
	if d.IPAddress == "" {
		return errors.New("dummy driver requires the --dummy-ip-address option")
	}

	return nil
}

func (d *Driver) PreCreateCheck() error {
	//if d.SSHKey != "" {
	//	if _, err := os.Stat(d.SSHKey); os.IsNotExist(err) {
	//		return fmt.Errorf("SSH key does not exist: %q", d.SSHKey)
	//	}

	//	// TODO: validate the key is a valid key
	//}

	log.Info("func PreCreateCheck()")
	return nil
}

func (d *Driver) Create() error {
	//if d.SSHKey == "" {
	//	log.Info("No SSH key specified. Assuming an existing key at the default location.")
	//} else {
	//	log.Info("Importing SSH key...")
//
	//	d.SSHKeyPath = d.ResolveStorePath(path.Base(d.SSHKey))
	//	if err := copySSHKey(d.SSHKey, d.SSHKeyPath); err != nil {
	//		return err
	//	}

	//	if err := copySSHKey(d.SSHKey+".pub", d.SSHKeyPath+".pub"); err != nil {
	//		log.Infof("Couldn't copy SSH public key : %s", err)
	//	}
	//}

	log.Info("func Create()")
	log.Debugf("IP: %s", d.IPAddress)

	return nil
}

func (d *Driver) GetURL() (string, error) {
	log.Info("func GetURL()")
	if err := drivers.MustBeRunning(d); err != nil {
		return "", err
	}

	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, strconv.Itoa(d.EnginePort))), nil
}

func (d *Driver) GetState() (state.State, error) {
	log.Info("func GetState()")
	//address := net.JoinHostPort(d.IPAddress, strconv.Itoa(d.SSHPort))

	//_, err := net.DialTimeout("tcp", address, defaultTimeout)
	//if err != nil {
	//	return state.Stopped, nil
	//}

	return state.Running, nil
}

func (d *Driver) Start() error {
	log.Info("func Start()")
	return errors.New("dummy driver does not support start")
}

func (d *Driver) Stop() error {
	log.Info("func Stop()")
	return errors.New("dummy driver does not support stop")
}

func (d *Driver) Restart() error {
	log.Info("func Restart()")
	//_, err := drivers.RunSSHCommandFromDriver(d, "sudo shutdown -r now")
	return errors.New("dummy driver does not support restart")
	//return err
}

func (d *Driver) Kill() error {
	log.Info("func Kill()")
	return errors.New("dummy driver does not support kill")
}

func (d *Driver) Remove() error {
	log.Info("func Remove()")
	return nil
}

func copySSHKey(src, dst string) error {
	log.Info("func copySSHKey()")
	if err := mcnutils.CopyFile(src, dst); err != nil {
		return fmt.Errorf("unable to copy ssh key: %s", err)
	}

	if err := os.Chmod(dst, 0600); err != nil {
		return fmt.Errorf("unable to set permissions on the ssh key: %s", err)
	}

	return nil
}
