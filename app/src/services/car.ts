import { Car } from "../types/models";
import api from "../utils/api";

export async function getCars(): Promise<Car[]> {
  return await api.get('cars')
}