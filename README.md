# 🔐 CLI Password Generator

Herramienta en Go para generar contraseñas seguras desde la línea de comandos.

## Uso

```bash
go run main.go --length 16 --symbols true --number 3
```

## Flags
- --length (int): longitud de la contraseña (default: 12)
- --symbols (bool): si se incluyen caracteres especiales
- --number (int): cantidad de contraseñas a generar
