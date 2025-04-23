// @ts-nocheck
import '@/App.scss'
import { RouterProvider } from '@tanstack/react-router'
import { router } from "@/routes/root";

function App() {

  return (
    <RouterProvider router={router} />
  )
}

export default App
