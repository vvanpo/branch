const merge = require("webpack-merge")

/*const clientConfig = {
    entry: "./client/index.ts",
}*/

const siteConfig = {
    mode: "production",
    entry: "./site/index.ts",
    output: {
        filename: "site.js",
        path: __dirname + "/dist",
    },
    module: {
        rules: [
            { test: /\.ts$/, use: "ts-loader" },
            { test: /\.html$/, use: "html-loader" },
            {
                // Copy files to dist/ without html-loader
                test: [__dirname + "/site/index.html"],
                use: {
                    loader: "file-loader",
                    options: { name: "[name].[ext]" },
                },
            },
            {
                test: /\.scss$/,
                use: ["css-loader", "sass-loader"],
            },
        ],
    },
}

const siteES5Config = {
    output: { filename: "site.es5.js" },
    module: { rules: [{ test: /\.ts$/, use: [
        { loader: "babel-loader", options: { presets: ["@babel/preset-env"] }},
        { loader: "ts-loader" },
    ]}]},
}

module.exports = (env, argv) => {
    if (argv.mode === "development") {
        siteConfig.devtool = "source-map"
    }

    return [siteConfig, merge.smart(siteConfig, siteES5Config)]
}
