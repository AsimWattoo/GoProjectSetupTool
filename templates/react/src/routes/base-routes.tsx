// @ts-nocheck
import rootRoute from "@/routes/root-route";
import { createRoute } from "@tanstack/react-router";
import Index from "@pages/index";

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Index,
});

const routes = [
    indexRoute
]

export default routes