import pycurl
from io import BytesIO

def MultiCurl():
  # Create a Curl object
c = pycurl.Curl()

# Set the URL
c.setopt(c.URL, 'http://example.com')

# Perform a first request
response = BytesIO()
c.setopt(c.WRITEFUNCTION, response.write)
c.perform()

# Perform a second request (reuses connection)
response2 = BytesIO()
c.setopt(c.WRITEFUNCTION, response2.write)
c.setopt(c.URL, 'http://example.com/some-other-page')
c.perform()

# Clean up
c.close()
