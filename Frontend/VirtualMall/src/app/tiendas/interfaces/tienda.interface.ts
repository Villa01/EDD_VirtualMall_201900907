export interface TiendasResponse {
    Datos: Dato[];
}

export interface Dato {
    Indice:        string;
    Departamentos: Departamento[];
}

export interface Departamento {
    Nombre:  string;
    Tiendas: Tienda[];
}

export interface Tienda {
    Nombre:       string;
    Descripcion:  string;
    Contacto:     string;
    Calificacion: number;
    Logo:         string;
}


export interface Producto {
    Nombre:      string;
    Codigo:      number;
    Descripcion: string;
    Precio:      number;
    Cantidad:    number;
    Imagen:      string;
}

export interface infoUsuario{
    DPI: number;
    password: string;
}

export interface RespuestaPassword {
    correcta: boolean;
    cuenta:   Cuenta;
}

export interface Cuenta {
    Dpi:      number;
    Nombre:   string;
    Correo:   string;
    Password: string;
    Usuario:  string;
}
