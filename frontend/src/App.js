import React, { useState } from 'react';
import UrlShortenerForm from './components/UrlShortenerForm';
import ShortUrlDisplay from './components/ShortUrlDisplay';
import Header from './components/Header';
import './styles/App.css';

function App() {
  const [shortUrl, setShortUrl] = useState('');
  const [error, setError] = useState('');

  const handleShortenSuccess = (shortCode) => {
    const fullShortUrl = `${shortCode}`;
    setShortUrl(fullShortUrl);
    setError('');
  };

  return (
    <div className="app">
      <Header />
      <div className="container">
        <UrlShortenerForm 
          onSuccess={handleShortenSuccess} 
          onError={(err) => setError(err)}
        />
        {shortUrl && <ShortUrlDisplay shortUrl={shortUrl} />}
        {error && <div className="error-message">{error}</div>}
      </div>
    </div>
  );
}

export default App;