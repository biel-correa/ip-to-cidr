const lastSubnetMask = () => {
  const urlParams = new URLSearchParams(window.location.search)
  const subnetMask = urlParams.get('subnetMask')
  return subnetMask != null && subnetMask <= 128 ? subnetMask : 32
}

const createSubnetMaskOptions = () => {
  const subnetMask = $('#subnetMask')
  for (let i = 0; i <= 128; i++) {
    subnetMask.append($('<option>', {
      value: i,
      text: `/${i}`
    }))
  }

  subnetMask.val(lastSubnetMask())
}

$(document).ready(() => {
  createSubnetMaskOptions()

  $('#useYourIp').click(() => {
    $(idInput).val(clientIp)
    $('#submit').click()
  })

  $('#copyResult').click(() => {
    navigator.clipboard.writeText(cidr)
  })
})