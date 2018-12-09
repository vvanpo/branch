/*const clientConfig = {
    entry: './client/index.ts',
}*/

const siteConfig = {
    entry: './site/index.ts',
    output: {
        filename: 'site.js',
        path: __dirname + '/dist',
    },
    module: {
        rules: [
            { test: /\.ts$/, use: 'ts-loader' },
            { test: /\.html$/, use: 'html-loader' },
            {
                // Copy the index.html to dist/ without html-loader
                test: [__dirname + '/site/index.html'],
                use: {
                    loader: 'file-loader',
                    options: { name: '[name].[ext]' },
                },
            },
            {
                test: /\.scss$/,
                use: ['css-loader', 'sass-loader'],
            }
        ],
    },
}

module.exports = (env, argv) => {
    return siteConfig
}
