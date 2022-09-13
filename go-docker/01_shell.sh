#!/usr/bin/env bash

echo "前段项目编译"
cd "G:\document\documents"
git pull origin main
npm run docs:build

