import { createBrowserRouter, RouterProvider } from "react-router-dom";
import CarPage from "./pages/car";
import CarsPage from "./pages/cars";

const router = createBrowserRouter([
  {
    path: "/cars",
    element: <CarsPage />,
  },
  {
    path: "/cars/:id",
    element: <CarPage />,
  },
]);

export default function AppRouter() {
  return <RouterProvider router={router} />;
}
