package repository

import (
	service2 "Parrillada/internal/comidas/service"
	"Parrillada/models/comensal"
	"Parrillada/models/comida"
	"Parrillada/utils"
	"reflect"
)

type Repository struct {
	Comensales     []comensal.IComensal
	comidasService service2.Service
}

func NewComensalesRepository(comensales []comensal.IComensal, comidasService service2.Service) Repository {
	return Repository{Comensales: comensales, comidasService: comidasService}
}

func (r *Repository) AgregarComensal(comensal comensal.IComensal) []comensal.IComensal {
	r.Comensales = append(r.Comensales, comensal)
	return r.Comensales
}

func (r *Repository) ObtenerComidasPreferidas(id string) []comida.IComida {

	for _, c := range r.Comensales {
		switch RetornarNombreStructHijoComensal(c) {
		case "*comensal.Vegetariano":
			v := c.(*comensal.Vegetariano)
			if v.ID == id {
				return v.ComidasFavoritas
			}

		case "*comensal.HambrePopular":
			hp := c.(*comensal.HambrePopular)
			if hp.ID == id {
				return hp.ComidasFavoritas
			}
		case "*comensal.PaladarFino":
			pf := c.(*comensal.PaladarFino)
			if pf.ID == id {
				return pf.ComidasFavoritas
			}
		}
	}

	return nil
}

/*
También se pide poder elegir un plato para un comensal - por ahora es cualquier plato que le guste.
Si no le gusta ningún plato, lanzar un error.
Si el plato existe, sacarlo de la cocina y hacer que el comensal lo coma.
*/

func (r *Repository) ElegirPlato(id string, comidasDispoRepo []comida.IComida) string {

	if len(comidasDispoRepo) < 1 {
		return "La cocina aun no tiene platos listos!"
	}

	for _, c := range r.Comensales {
		switch RetornarNombreStructHijoComensal(c) {
		case "*comensal.Vegetariano":
			v := c.(*comensal.Vegetariano)
			if v.ID == id {
				if len(v.ComidasFavoritas) < 1 {
					return "El comensal no tiene platos favoritos"
				}

				//Arbitrariamente elijo el primero de los favoritos
				plato := v.ComidasFavoritas[0]
				platoDecodificado := utils.RetornarNombreStructHijoComida(plato)
				//Decodifico el plato y veo si le gusta
				for _, comidaDisponibles := range comidasDispoRepo {
					comidaDisponibleDecodificada := utils.RetornarNombreStructHijoComida(comidaDisponibles)

					if comidaDisponibleDecodificada == platoDecodificado {
						switch comidaDisponibleDecodificada {
						case "*comida.Provoleta":
							pd := comidaDisponibles.(*comida.Provoleta)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "provoleta consumida"
						case "*comida.HamburguesaVegana":
							pd := comidaDisponibles.(*comida.HamburguesaVegana)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa vegana consumida"
						case "*comida.HamburguesaDeCarne":
							pd := comidaDisponibles.(*comida.HamburguesaDeCarne)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa de carne consumida"
						case "*comida.Parrillada":
							pd := comidaDisponibles.(*comida.Parrillada)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Parrillada consumida"
						}
					}

				}

			}

		case "*comensal.HambrePopular":
			v := c.(*comensal.HambrePopular)
			if v.ID == id {
				if len(v.ComidasFavoritas) < 1 {
					return "El comensal no tiene platos favoritos"
				}

				//Arbitrariamente elijo el primero de los favoritos
				plato := v.ComidasFavoritas[0]
				platoDecodificado := utils.RetornarNombreStructHijoComida(plato)
				//Decodifico el plato y veo si le gusta
				for _, comidaDisponibles := range comidasDispoRepo {
					comidaDisponibleDecodificada := utils.RetornarNombreStructHijoComida(comidaDisponibles)

					if comidaDisponibleDecodificada == platoDecodificado {
						switch comidaDisponibleDecodificada {
						case "*comida.Provoleta":
							pd := comidaDisponibles.(*comida.Provoleta)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "provoleta consumida"
						case "*comida.HamburguesaVegana":
							pd := comidaDisponibles.(*comida.HamburguesaVegana)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa vegana consumida"
						case "*comida.HamburguesaDeCarne":
							pd := comidaDisponibles.(*comida.HamburguesaDeCarne)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa de carne consumida"
						case "*comida.Parrillada":
							pd := comidaDisponibles.(*comida.Parrillada)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Parrillada consumida"
						}
					}

				}

			}
		case "*comensal.PaladarFino":
			v := c.(*comensal.PaladarFino)
			if v.ID == id {
				if len(v.ComidasFavoritas) < 1 {
					return "El comensal no tiene platos favoritos"
				}

				//Arbitrariamente elijo el primero de los favoritos
				plato := v.ComidasFavoritas[0]
				platoDecodificado := utils.RetornarNombreStructHijoComida(plato)
				//Decodifico el plato y veo si le gusta
				for _, comidaDisponibles := range comidasDispoRepo {
					comidaDisponibleDecodificada := utils.RetornarNombreStructHijoComida(comidaDisponibles)

					if comidaDisponibleDecodificada == platoDecodificado {
						switch comidaDisponibleDecodificada {
						case "*comida.Provoleta":
							pd := comidaDisponibles.(*comida.Provoleta)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "provoleta consumida"
						case "*comida.HamburguesaVegana":
							pd := comidaDisponibles.(*comida.HamburguesaVegana)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa vegana consumida"
						case "*comida.HamburguesaDeCarne":
							pd := comidaDisponibles.(*comida.HamburguesaDeCarne)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Hamburguesa de carne consumida"
						case "*comida.Parrillada":
							pd := comidaDisponibles.(*comida.Parrillada)
							v.Comer(pd)
							r.comidasService.EliminarPlato(pd.ID)
							return "Parrillada consumida"
						}
					}

				}

			}
		}
	}
	return "No se pudo encontrar un plato aconsejable para el comensal"
}

func (r *Repository) LeAgradaPlato(plato comida.IComida, id string) []comida.IComida {
	for _, c := range r.Comensales {
		switch RetornarNombreStructHijoComensal(c) {
		case "*comensal.Vegetariano":
			v := c.(*comensal.Vegetariano)
			if v.ID == id {
				if v.LeAgradaComida(plato) {
					return v.ComidasFavoritas
				}
				return v.ComidasFavoritas
			}

		case "*comensal.HambrePopular":
			hp := c.(*comensal.HambrePopular)
			if hp.ID == id {
				if hp.LeAgradaComida(plato) {
					return hp.ComidasFavoritas
				}
				return hp.ComidasFavoritas
			}
		case "*comensal.PaladarFino":
			pf := c.(*comensal.PaladarFino)
			if pf.ID == id {
				if pf.LeAgradaComida(plato) {
					return pf.ComidasFavoritas
				}
				return pf.ComidasFavoritas
			}
		}
	}
	return nil
}

//TODO reubicar bien esto para evitar impotacion ciclica
func RetornarNombreStructHijoComensal(comensal comensal.IComensal) string {

	childStruct := reflect.TypeOf(comensal)

	return childStruct.String()
}
