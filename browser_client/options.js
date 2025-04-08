document.getElementById('save-settings').addEventListener('click', () => {
  const torAddress = document.getElementById('tor-address').value;
  const ipv0xServer = document.getElementById('ipv0x-server').value;

  chrome.storage.sync.set({ torAddress, ipv0xServer }, () => {
    alert("Settings Saved");
  });
});

// Load saved settings
chrome.storage.sync.get(['torAddress', 'ipv0xServer'], (data) => {
  document.getElementById('tor-address').value = data.torAddress || '';
  document.getElementById('ipv0x-server').value = data.ipv0xServer || '';
});
