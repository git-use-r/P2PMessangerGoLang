## *ВНИМАНИЕ!!! ОТСУТСТВИЕ ЗАЩИТЫ ОТ MITM (Man In The Middle) АТАКИ*

### P2PMessangerGoLang
  **P2PMessangerGoLang** - Это P2P мессенджер на базе голого TLS соеденения. Весь проект написан на Go с использоание GUI фреймворка Wails. Frontend часть сипользует чистый JavaScript, CSS и HTML. 
  *Мессенджер не имеет защиты от MITM*

  ### Для пользовтелей
   Как использовать? 
   * Перейдите в раздел Releases И скачайте последнюю версию мессенджера в .exe формате. Или же можете собрать сами .exe из исходников.

 ### Для Разработчиков
 Требования
 * Убедитесь что у вас устновлен Go версии "1.23.0".
 * Убедитесь что у вас устновлен Wails CLI.
 * Убедитесь что у вас установлен Node.js.

 * Клонирование репозитория 
```bash
git clone https://github.com/git-use-r/P2PMessangerGoLang
cd P2PMessangerGoLang
```
 * Запуск приложения с hot-reload
```bash
wails dev
```
* Сборка приложения
```bash
wails build
```
