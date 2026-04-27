### Hexlet tests and linter status:

[![Actions Status](https://github.com/AlinaMavlekaeva/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/AlinaMavlekaeva/go-from-scratch-project-242/actions)

## hexlet-path-size

### About

hexlet-path-size - CLI утилита, которая выводит на экран информацию о размере файла или директории в виде `размер   путь`
Утилита имеет 3 основных флага и флаг помощи.

**Доступные флаги:**
1. **--human (-H)** - отвечает за вывод информации в человекочитаемом формате 
2. **--all (-a)** - отвечает за учет в выводе скрытых файлов и директорий (скрытыми являются файлы, имена которых начинаются с ".")
3. **--recursive (-r)** - отвечает за учет файлов и директорий на всех уровнях вложенности;
4. **--help (-h)** - выводит на экран описание утилиты. 

### Install

Для установки утилиты нужно выполнить следующую команду:
```
go install github.com/AlinaMavlekaeva/go-from-scratch-project-242@latest
```

Проверить работу утилиты можно запустив команду:
```
hexlet-path-size -h
``` 
В результате должна появиться информационная справка вида:
```
NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options]

GLOBAL OPTIONS:
   --human, -H      human-readable sizes (auto-select unit)
   --all, -a        include hidden files and directories
   --recursive, -r  recursive size of directories
   --help, -h       show help
   ```

 ### Examples

Для получения информации о размере файла или директории необходимо передать в утилиту относительный путь. 

**Работа утилиты без передачи флагов. Размер директории выводится в байтах.**
Команда:   
 ```
 hexlet-path-size test
 ```            
Вывод:      
```
1559B   test
```

**Определение размера директории в человекочитаемом формате (флаг --human)**
Команда:
```
hexlet-path-size test -H
```        
Вывод: 
```
1.5KB   test
```

**Определение размера директории с учетом скрытых файлов (флаг --all)**
Команда: 
```
hexlet-path-size test -H -a
```      
Вывод: 
```
3.1KB   test
```
**Определение размера директории с учетом всех уровней вложенности (флаг --recursive)**
Команда: 
```
hexlet-path-size test -H -r
```    
Вывод: 
```
3.1KB   test
```
**Определение размера директории с учетом скрытых файлов на всех уровнях вложенности**
Команда: 
```
hexlet-path-size test -H -a -r
```      
Вывод: 
```
6.1KB   test
```


Asciinema link https://asciinema.org/a/GpujsORGqrXSWRLU
