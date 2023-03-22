# Getting started

Проект содержит следующие директории:

- `app` - директория с сервисом, который получает IP и записывает его в БД
- `charts/app` - helm chart сервиса
- `charts/postgres` - helm chart postgres'а
- `playbook` - плейбук, который все это устанавливает
- `minikube.sh` - скрипт для запуска плейбука в minikube

# Usage example

1. Клонируем репо

```
git clone https://github.com/loktev-d/kts-test
```

2. Перед тем, как запускать плейбук, необходимо поменять значения для своей среды в следующих файлах:

- `charts/postgres/values.yaml`
- `charts/app/values.yaml`
- `playbook/hosts.ini`

3. Переходим в `playbook/`

```
cd playbook/
```

4. Устанавливаем зависимости из galaxy

```
ansible-galaxy install -r requirements.yaml
```

5. Запускаем плейбук

```
ansible-playbook --skip-tags "prereq" install.yaml
```

Блок с тегом `prereq` используется для установки необходимых модулей python и helm, мне они были нужны для того, чтобы плейбук работал в minikube. Если на хосте уже есть все необходимое, его можно пропустить.

6. Ждем пока поднимутся поды

```
kubectl get pods -w
```

7. Готово!

По умолчанию ингресс сервиса получения IP установлен на `/app`, я не ставил хост, чтобы не нужно было возиться с DNS.

В зависимости от среды пункты выше могут отличаться, я тестировал в `minikube`.
