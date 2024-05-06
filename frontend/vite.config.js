export default {
  server: {
    proxy: {
      "/chatroom": {
        target: "ws://localhost:8080/ws",
        ws: true,
      },
    },
  },
};
