# FileHandler
A CLI file handler in Go

### Comandos Disponíveis
**help**: Exibe uma mensagem de ajuda mostrando uma lista de comandos disponíveis e suas descrições.
```help```

**read**: Lê o conteúdo de um arquivo e o exibe na saída padrão.
```read example.txt```

**show**: Lista os arquivos e diretórios no diretório especificado ou no diretório atual se nenhum diretório for fornecido.
```show /path/to/directory```

**info**: Exibe informações detalhadas sobre um arquivo, incluindo nome, tamanho, modo e data de modificação.
```info example.txt```

**exist**: Verifica se um arquivo ou diretório existe.
```exist example.txt```

**start**: Cria um novo arquivo.
```start newfile.txt```

**startDir**: Cria um novo diretório.
```startDir newdirectory```

**copy**: Copia um arquivo para outro local.
```copy file1.txt /path/to/destination/file1.txt```

**add**: Adiciona texto a um arquivo existente.
```add "Novo conteúdo" example.txt```

**switch**: Renomeia um arquivo ou diretório.
```switch oldfile.txt newfile.txt```

**move**: Move um arquivo de um local para outro.
```move file.txt /path/to/destination```

**goto**: Muda para o diretório especificado.
```goto /path/to/directory```

**del**: Deleta um arquivo ou diretório.
```del file.txt```
