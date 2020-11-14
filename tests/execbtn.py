import sys, getopt

from selenium import webdriver
from selenium.webdriver.firefox.firefox_binary import FirefoxBinary
from selenium.webdriver import FirefoxOptions

def main(argv):
    # Handle command line.
    firefoxpath = ''

    try:
        opts, args = getopt.getopt(argv, "p:h", ["path="])
    except getopt.GetoptError:
        print('execbtn.py -p <path to Firefox binary>')

        sys.exit(2)
    
    for opt, arg in opts:
        if opt in ("-p", "--path"):
            firefoxpath = arg
        elif opt == "-h":
            print('execbtn.py -p <path to Firefox binary>')
            
            sys.exit(0)

    # Run firefox in headless mode.
    opts = FirefoxOptions()
    opts.add_argument("--headless")

    # Use custom Firefox location.
    binary = FirefoxBinary(firefoxpath)

    # Launch Firefox driver setting custom binary path and firefox options.
    driver = webdriver.Firefox(firefox_binary=binary, firefox_options=opts)

    # Get our website.
    driver.get("http://127.0.0.1:8808")

    # Let's click our button via JavaScript.
    driver.execute_script("document.getElementById('btn').click()")

    # Get contents of ID 'fill'.
    fill = driver.find_element_by_id('fill')
    html = fill.get_attribute('innerHTML')

    driver.close()

    if len(html) == 0:
        print("Button not successful.")
        sys.exit(1)

if __name__ == "__main__":
    main(sys.argv[1:])