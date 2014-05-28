package main

// {
//   "name": "dtesting",
//   "image": "charliek/docker-testing:green",
//   "stdin": false,
//   "stdout": true,
//   "stderr": true,
//   "ports": [
//     {
//       "protocol": "tcp",
//       "hostPort": 9090,
//       "containerPort": 9090,
//     }
//   ],
//   "env": {
//     "FOO": "BAR",
//     "VAR2": "value"
//   },
//   "command": [
//     "/bin/echo",
//     "'12345' '6789'"
//   ]
// }

type PortMapping struct {
	Protocol      string
	HostPort      uint16
	ContainerPort uint16
}

type UpdocApp struct {
	Name    string
	image   string
	Stdin   bool
	Stdout  bool
	Stderr  bool
	Ports   []PortMapping
	Env     map[string]string
	Command []string
}
