import React from 'react';
import DistributorForm from '@/components/DistributorForm';

const App = () => {
  return (
    <div className="flex flex-col min-h-screen bg-slate-50">
      <header className="py-8 bg-white shadow-sm border-b border-slate-200">
        <div className="container mx-auto px-4 max-w-3xl">
          <h1 className="text-4xl font-bold text-slate-900 text-center tracking-tight">
            Distributor Management
          </h1>
          <p className="mt-3 text-lg text-slate-600 text-center max-w-xl mx-auto">
            Create and manage your distributors and their regions efficiently
          </p>
        </div>
      </header>

      <main className="flex-1">
        <div className="container mx-auto py-12 px-4 max-w-3xl">
          <div className="max-w-2xl mx-auto">
            <DistributorForm />
          </div>
        </div>
      </main>

      <footer className="py-8 bg-white border-t border-slate-200">
        <div className="container mx-auto px-4 max-w-3xl">
          <div className="flex flex-col items-center justify-center space-y-4">
            <p className="text-sm text-slate-500">
              © 2024 Qube Cinema. All rights reserved.
            </p>
            <div className="flex items-center space-x-4">
              <a href="#" className="text-sm text-slate-500 hover:text-slate-700">Terms</a>
              <a href="#" className="text-sm text-slate-500 hover:text-slate-700">Privacy</a>
              <a href="#" className="text-sm text-slate-500 hover:text-slate-700">Contact</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default App;