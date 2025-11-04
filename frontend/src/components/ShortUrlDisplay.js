import React from 'react';

function ShortUrlDisplay({ shortUrl }) {
  const copyToClipboard = async () => {
    if (navigator.clipboard && window.isSecureContext) {
      try {
        await navigator.clipboard.writeText(shortUrl);
        alert('Copied to clipboard!');
      } catch (err) {
        console.error('Clipboard API failed:', err);
      }
    } else {
      const textArea = document.createElement('textarea');
      textArea.value = shortUrl;
      textArea.style.position = 'fixed';
      document.body.appendChild(textArea);
      textArea.focus();
      textArea.select();
      try {
        const successful = document.execCommand('copy');
        if (successful) {
          alert('Copied to clipboard!');
        } else {
          console.error('Fallback: Copy command failed.');
        }
      } catch (err) {
        console.error('Fallback: Unable to copy', err);
      }
      document.body.removeChild(textArea);
    }
  };

  return (
    <div className="short-url-container">
      <h3>Your shortened URL:</h3>
      <div className="short-url">
        <a href={shortUrl} target="_blank" rel="noopener noreferrer">
          {shortUrl}
        </a>
        <button onClick={copyToClipboard}>Copy</button>
      </div>
    </div>
  );
}

export default ShortUrlDisplay;