FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV VITE_BASE_URL="https://m.mixkomii.com"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

COPY . /app
WORKDIR /app

RUN pnpm install

RUN pnpm build

FROM node:20-slim AS production
WORKDIR /app
COPY --from=base /app/dist /app/dist
COPY --from=base /app/package.json /app/package.json
COPY --from=base /app/node_modules /app/node_modules
EXPOSE 3000
ENTRYPOINT ["node", "dist/server.js"]