package exception

type NotFoundException struct {
	Error string
}

func NewNotFoundException(message string) NotFoundException {
	return NotFoundException{Error: message}
}

type UnprocessableEntityException struct {
	Error string
}

func NewUnprocessableEntityException(message string) UnprocessableEntityException {
	return UnprocessableEntityException{Error: message}
}

type CredentialException struct {
	Error string
}

func NewCredentialException(message string) CredentialException {
	return CredentialException{Error: message}
}

type BadRequestException struct {
	Error string
}

func NewBadRequestException(message string) BadRequestException {
	return BadRequestException{Error: message}
}
