package books

import "errors"

var ErrTitleRequired = errors.New("title required")
var ErrPublisherRequired = errors.New("publisher required")
var ErrPagesRequired = errors.New("pages required")
var ErrLanguageRequired = errors.New("language required")
var ErrEditionRequired = errors.New("edition required")
var ErrYearRequired = errors.New("year required")
var ErrISBNRequired = errors.New("isbn required")
var ErrOwnerRequired = errors.New("owner required")
