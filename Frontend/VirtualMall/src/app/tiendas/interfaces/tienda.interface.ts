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

interface HTMLInputEvent extends Event {
    target: HTMLInputElement & EventTarget;
}