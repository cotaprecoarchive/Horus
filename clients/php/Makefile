PHPUNIT="./vendor/bin/phpunit"

tests:
	@$(PHPUNIT) --colors tests/.

check-cs:
	@./vendor/bin/phpcs --standard=PSR2 src tests

.PHONY: tests check-cs
