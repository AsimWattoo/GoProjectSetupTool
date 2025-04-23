// @ts-nocheck
import { createRouter } from "@tanstack/react-router";
import rootRoute from "@/routes/root-route";
import baseRoutes from "@/routes/base-routes";

const routeTree = rootRoute.addChildren([
    ...baseRoutes
])

export const router = createRouter({
   routeTree
})