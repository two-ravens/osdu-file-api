// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AppError struct {
	Code    *int
	Reason  *string
	Message *string
}

func (o *AppError) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *AppError) GetReason() *string {
	if o == nil {
		return nil
	}
	return o.Reason
}

func (o *AppError) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
