// @ts-nocheck
import { defineConfig } from 'vitest/config'
import { resolve } from "path";

// https://vite.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            "@": resolve(__dirname, "src"),
        }
    },
    test: {
        environment: "jsdom",
        include: ["test/**/*.spec.ts"],
        setupFiles: ["./test/vitest.setup.ts"],
        globals: true,
        testTimeout: 5000,
        reporters: "verbose",
    }
})
