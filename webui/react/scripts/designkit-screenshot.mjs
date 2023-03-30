import fs from 'fs/promises';
import path from 'path';

import { chromium } from 'playwright';
import { createServer } from 'vite';

const THEMES = ['light', 'dark'];

const label = process.argv[2] || '';
const screenPath = path.resolve(process.cwd(), ...['screenshots', label].filter((c) => c));

// start dev server
const devServer = await createServer({
  mode: 'test', // disables typescript checking
  server: {
    open: false,
    port: 3456,
  },
});
await devServer.listen();

// start chrome playwright
const browser = await chromium.launch();
const page = await browser.newPage();
await page.goto('http:localhost:3456/src/designkit-standalone/');
// take screenshots of each section
const sections = await page.locator('article > section').all();
for (const theme of THEMES) {
  const themePath = path.resolve(screenPath, theme);
  await fs.mkdir(themePath, { recursive: true });
  await page.emulateMedia({ colorScheme: theme });
  for (const section of sections) {
    const header = await section.locator('h3').innerText();
    // playwright hangs if height is a non-int
    const height = Math.ceil((await section.boundingBox()).height);
    await page.setViewportSize({ height, width: 1280 });
    await section.screenshot({
      animations: 'disabled',
      path: path.resolve(themePath, `${header}.png`),
    });
  }
}

// clean up
await browser.close();
await devServer.close();
