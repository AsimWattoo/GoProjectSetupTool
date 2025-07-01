// @ts-nocheck
import { defineConfig } from 'vite'
import { resolve } from "path";
import dts from "vite-plugin-dts";

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        dts({
            insertTypesEntry: true,
            rollupTypes: true,
            include: ["src"],
            exclude: [
                "**/*.test.*",
                "**/*.spec.*",
            ],
            tsconfigPath: "./tsconfig.lib.json",
            outDir: "dist",
            pathsToAliases: true,
            staticImport: true,
            clearPureImport: true
        }),
    ],
    build: {
        lib: {
            entry: resolve(__dirname, 'src/index.ts'),
            name: '<PACKAGE_NAME>',
            fileName: (format) => {
                if(format === "es") return "<PACKAGE_NAME>.js";
                else if (format === "umd") return "<PACKAGE_NAME>.umd.cjs";
                return `<PACKAGE_NAME>.${format}.js`
            },
            formats: ['es', 'umd', 'cjs']
        },
        emptyOutDir: true,
        sourcemap: true,
    },
    resolve: {
        alias: {
            "@": resolve(__dirname, "src"),
        }
    },
})
