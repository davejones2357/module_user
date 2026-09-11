module module_user

go 1.26.2

require module_provider v0.0.0

replace module_provider => ../module_provider //use local path so that we can test the module_provider changes without publishing it to a remote repository

