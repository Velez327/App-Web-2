package cafeteria

import (
		//"fmt"
		"errors"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Telefono int
	Email string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
	Detalles  string
}

type Pedido struct {
	ID         int
	Cliente []Cliente 
	Producto []Producto
	Cantidad   int
	Totalcantidad    float64
	Fecha      string
	Ruc	   string
}

var (
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente = errors.New("stock insuficiente")
	ErrSaldoInsuficiente = errors.New("saldo insuficiente")
)

type Repository interface {
	GuardarC(c Cliente) error
	//AgregarC(c Cliente) error
	BuscarPorIDC(id int) (Cliente, error)
	ListarC() []Cliente
	GuardarP(p Producto) error
	//AgregarP(p Producto) error
	BuscarPorIDP(id int) (Producto, error)
	ListarP() []Producto

}

type RepoMemoria struct{
	clientes []Cliente
	productos []Producto
	pedidos []Pedido
}

/*func ResivirPedido(p Pedido) error {
	return nil
	
}*/
func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{
		clientes: []Cliente{},
		productos: []Producto{},
		pedidos:   []Pedido{},
	}
}
func (r *RepoMemoria) GuardarC(c Cliente) error{
	r.clientes = append(r.clientes, c)
	return nil
}
func (r *RepoMemoria) GuardarP(p Producto) error{
	r.productos = append(r.productos, p)
	return nil
}
func (r *RepoMemoria) BuscarPorIDC(id int) (Cliente, error){
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}
func (r *RepoMemoria) BuscarPorIDP(id int) (Producto, error){
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}
func (r *RepoMemoria) ListarC() []Cliente{
	return r.clientes
}
func (r *RepoMemoria) ListarP() []Producto{
	return r.productos
}

var _ Repository = (*RepoMemoria)(nil)
