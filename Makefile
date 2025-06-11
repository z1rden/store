build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	docker-compose up --force-recreate --build

run:
	docker-compose up -d --force-recreate

stop:
	docker-compose down
	docker rmi cart loms

# Добавление удаленного репозитория cart:
# 1. Сначала добавляется удаленный репозиторий cart
# 2. Затем git fetch загружает актуальные данные из удалённого репозитория в локальное хранилище, но не изменяет файлы в рабочей директории. (альтернатива - git pull).
.git-add-cart:
	git remote add cart https://github.com/z1rden/cart.git
	git fetch cart

# Добавление удаленного репозитория loms:
# 1. Сначала добавляется удаленный репозиторий loms
# 2. Затем git fetch загружает актуальные данные из удалённого репозитория в локальное хранилище, но не изменяет файлы в рабочей директории. (альтернатива - git pull).
.git-add-loms:
	git remote add loms https://github.com/z1rden/loms.git
	git fetch cart

# Влить ветку внешнего репозитория в папку `cart`
# --squash объединяет всю историю внешнего репозитория в один коммит
.git-subtree-cart:
	git subtree add --prefix=cart cart main --squash

# Влить ветку внешнего репозитория в папку `cart`
# --squash объединяет всю историю внешнего репозитория в один коммит
.git-subtree-loms:
	git subtree add --prefix=loms loms main --squash

# Для обновления в рабочей директории сервиса cart
.git-update-cart:
	git fetch cart
	git subtree pull --prefix=cart cart main --squash

# Для обновления в рабочей директории сервиса loms
.git-update-loms:
	git fetch loms
	git subtree pull --prefix=loms loms main --squash