// scripts/check-package-manager.ts
const allowedPackageManager = 'pnpm';

if (process.env.npm_execpath) {
  const execPath = process.env.npm_execpath.split('/');
  const usedPackageManager = execPath[execPath.length - 1];

  if (usedPackageManager !== allowedPackageManager) {
    console.error(`Please use ${allowedPackageManager} instead of ${usedPackageManager} to install dependencies.`);
    process.exit(1);
  }
}

