const config = {
  title: "Gofast",
  tagline: "Go-first server-rendered apps with a tiny SPA layer",
  favicon: "img/favicon.ico",
  url: "https://example.com",
  baseUrl: "/",
  organizationName: "ezeki",
  projectName: "gofast",
  onBrokenLinks: "throw",
  markdown: {
    hooks: {
      onBrokenMarkdownLinks: "warn",
    },
  },
  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: require.resolve("./sidebars.js"),
          routeBasePath: "/",
        },
        blog: false,
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      },
    ],
  ],
  themeConfig: {
    navbar: {
      title: "Gofast",
      items: [
        { type: "docSidebar", sidebarId: "docs", position: "left", label: "Docs" },
        { to: "/download", label: "Download", position: "left" },
      ],
    },
  },
};

module.exports = config;
