document.getElementById('enable').addEventListener('click', () => {
  chrome.storage.sync.set({ ipv0xEnabled: true });
  alert("IPv0x Enabled");
});

document.getElementById('disable').addEventListener('click', () => {
  chrome.storage.sync.set({ ipv0xEnabled: false });
  alert("IPv0x Disabled");
});

document.getElementById('server-config').addEventListener('click', () => {
  chrome.runtime.openOptionsPage();
});
